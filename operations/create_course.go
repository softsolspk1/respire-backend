package operations

import (
	"net/http"
	"time"

	"fr_book_api/models"

	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/zap"
)

// CreateCourse creates a new course (admin only)
func CreateCourse(sugar string, mongoDb *mongo.Database, logger *zap.Logger) http.Handler {
	oLog := logger.With(zap.String("op", "createCourse"))
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		v := models.NewValidator(r).Secret(sugar)

		userId := v.Token("user_id").Int()
		// Check if user is an admin (assuming role is stored in token)
		isAdmin := v.Token("admin").Bool()

		log := oLog.With(zap.String("ip", r.Header.Get("X-Real-IP")))
		
		if !v.Valid() {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		// Only admin can create courses
		if !isAdmin {
			log.Debug("Non-admin attempted to create course", zap.Int("user_id", userId))
			w.WriteHeader(http.StatusForbidden)
			return
		}

		course := v.CourseFromBody()
		if course == nil {
			log.Debug("Invalid course data", zap.Int("user_id", userId))
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		// Generate ID for the course
		idm := models.IdMgr{DB: mongoDb}
		courseId, err := idm.NextId("courses")
		if err != nil {
			log.Error("Failed to generate ID", zap.Error(err))
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		// Set course metadata
		course.Id = courseId
		course.CreatedAt = time.Now()
		course.CreatedBy = userId

		// Insert course into database
		_, err = mongoDb.Collection("courses").InsertOne(r.Context(), course)
		if err != nil {
			log.Error("Failed to insert course", zap.Error(err))
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		// Return success response
		JSON(&models.CourseResponse{
			Code:   200,
			Result: *course,
		}, w)
	})
}