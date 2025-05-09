package models

import (
	"encoding/json"
	"net/http"
	"net/url"
	"strconv"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
)

var VERSION = "1.0.0"
var BUILD = 1

type Values struct {
	v        *Validator
	d        []string
	def      *string
	name     string
	optional bool
}

func (v *Values) Def(d string) *Values {
	v.def = &d
	return v
}

func (v *Values) Optional() *Values {
	v.optional = true
	return v
}

func (v *Values) Int() int {
	s := v.String()
	if s == "" && v.optional {
		return 0
	}
	ret, err := strconv.Atoi(s)
	if err != nil {
		v.v.Error(v.name, err.Error())
		return 0
	}
	return ret
}

func (v *Values) Bool() bool {
	s := v.String()
	if s == "" && v.optional {
		return false
	}
	ret, err := strconv.ParseBool(s)
	if err != nil {
		v.v.Error(v.name, err.Error())
		return false
	}
	return ret
}

func (v *Values) String() string {
	if len(v.d) == 0 {
		if v.def == nil {
			if !v.optional {
				v.v.Error(v.name, "Missing Value for "+v.name)
			}
			return ""
		}
		return *v.def
	}
	return v.d[0]
}

func (v *Values) StringArray() []string {
	return v.d
}

type Validator struct {
	r       *http.Request
	m       map[string]string // Mux Vars
	q       url.Values        // Query Vars
	j       jwt.MapClaims     // JWT Claims
	values  []string
	errors  map[string]string
	_secret string
}

func NewValidator(r *http.Request) *Validator {
	return &Validator{r: r, m: mux.Vars(r), q: r.URL.Query(), errors: make(map[string]string)}
}

func (v *Validator) Secret(secret string) *Validator {
	v._secret = secret
	return v
}

func (v *Validator) Error(name string, msg string) {
	if v.errors == nil {
		v.errors = make(map[string]string)
	}
	v.errors[name] = msg
}

func (v *Validator) Path(name string) *Values {
	if v.m == nil {
		v.m = mux.Vars(v.r)
	}
	val, ok := v.m[name]
	if !ok {
		return v.nilValues(name)
	}
	return &Values{v: v, d: []string{val}, name: name}
}

func (v *Validator) nilValues(name string) *Values {
	return &Values{v: v, d: []string{}, name: name}
}

func (v *Validator) jwtSugarFn(token *jwt.Token) (interface{}, error) {
	return []byte(v._secret), nil
}

func (v *Validator) Token(name string) *Values {
	if v.j != nil {
		val, ok := v.j[name]
		if ok {
			switch val.(type) {
			case float64:
				// JSON Numbers end up as float64 in golang by default, convert to string
				strVal := strconv.FormatFloat(val.(float64), 'f', -1, 64)
				return &Values{v: v, d: []string{strVal}, name: name}
			case string:
				return &Values{v: v, d: []string{val.(string)}, name: name}
			case bool:
				return &Values{v: v, d: []string{strconv.FormatBool(val.(bool))}, name: name}
			default:
				jsonVal, err := json.Marshal(val)
				if err != nil {
					return v.nilValues(name)
				}
				return &Values{v: v, d: []string{string(jsonVal)}, name: name}
			}
		}
	}

	authHeader := v.Header("jwt").String()
	if authHeader == "" {
		return v.nilValues(name)
	}

	token, err := jwt.Parse(authHeader, v.jwtSugarFn)
	if err != nil {
		v.Error("token", err.Error())
		return v.nilValues(name)
	}

	if !token.Valid {
		v.Error("token", "Invalid Token")
		return v.nilValues(name)
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		v.Error("token", "Invalid Claims")
		return v.nilValues(name)
	}

	v.j = claims
	val, ok := v.j[name]
	if !ok {
		return v.nilValues(name)
	}

	switch val.(type) {
	case float64:
		// JSON Numbers end up as float64 in golang by default, convert to string
		strVal := strconv.FormatFloat(val.(float64), 'f', -1, 64)
		return &Values{v: v, d: []string{strVal}, name: name}
	case string:
		return &Values{v: v, d: []string{val.(string)}, name: name}
	case bool:
		return &Values{v: v, d: []string{strconv.FormatBool(val.(bool))}, name: name}
	default:
		jsonVal, err := json.Marshal(val)
		if err != nil {
			return v.nilValues(name)
		}
		return &Values{v: v, d: []string{string(jsonVal)}, name: name}
	}
}

func (v *Validator) Query(name string) *Values {
	if v.q == nil {
		v.q = v.r.URL.Query()
	}
	return &Values{v: v, d: v.q[name], name: name}
}

func (v *Validator) Header(name string) *Values {
	return &Values{v: v, d: v.r.Header[name], name: name}
}

func (v *Validator) Valid() bool {
	return len(v.errors) == 0
}

func (v *Validator) Errors() map[string]string {
	return v.errors
}