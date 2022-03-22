package Controller

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"
	//"github.com/gorilla/mux"
)

//GetAllSchedule
func GetAllSchedule(w http.ResponseWriter, r *http.Request, ren render.Render) {
	db := connect()
	defer db.Close()

	query := "SELECT schedule.scheduleID, schedule.scheduleDateTime, " +
		"patient.patientID, patient.patientName, patient.patientAge, patient.patientAddress, patient.patientGender, " +
		"medical_staff.medstaffID, medical_staff.medstaffName, medical_staff.medstaffSpecialty FROM schedule " +
		"JOIN medical_staff ON schedule.medstaffID = medical_staff.medstaffID " +
		"JOIN patient ON schedule.patientID = patient.patientID "

	rows, errQuery := db.Query(query)
	if errQuery != nil {
		log.Println(errQuery)
	}

	var scheduleDetail ScheduleDetail
	var scheduleDetails []ScheduleDetail
	if errQuery == nil {
		for rows.Next() {
			if err := rows.Scan(&scheduleDetail.ID, &scheduleDetail.DateTime, &scheduleDetail.Patient.ID,
				&scheduleDetail.Patient.Name, &scheduleDetail.Patient.Age, &scheduleDetail.Patient.Address, &scheduleDetail.Patient.Gender,
				&scheduleDetail.Medstaff.ID, &scheduleDetail.Medstaff.Name, &scheduleDetail.Medstaff.Specialty); err != nil {
				log.Fatal(err.Error())
			} else {
				scheduleDetails = append(scheduleDetails, scheduleDetail)
			}
		}
		ren.HTML(http.StatusOK, "schedule_details", scheduleDetails)
	}
	var response ScheduleDetailResponse
	if len(scheduleDetails) > 0 {
		response.Status = 200
		response.Message = "Success"
		response.Data = scheduleDetails
	} else {
		response.Status = 400
		response.Message = "Fail to retrieve schedule! " + errQuery.Error()
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

//GetMedStaffSchedule
func GetMedStaffSchedule(w http.ResponseWriter, r *http.Request, ren render.Render) {
	db := connect()
	defer db.Close()

	query := "SELECT * FROM schedule"

	medstaffID := r.URL.Query()["medstaffid"]
	if medstaffID != nil {
		query += " WHERE medstaffID = '" + medstaffID[0] + "'"
	}
	rows, err := db.Query(query)
	if err != nil {
		log.Println(err)
	}

	var schedule Schedule
	var schedules []Schedule
	for rows.Next() {
		if err := rows.Scan(&schedule.ID, &schedule.DateTime, &schedule.PatientID, &schedule.MedstaffID); err != nil {
			log.Fatal(err.Error())
		} else {
			schedules = append(schedules, schedule)
		}
	}
	ren.HTML(http.StatusOK, "schedules", schedules)
	var response ScheduleResponse
	if err == nil {
		response.Status = 200
		response.Message = "Success"
		response.Data = schedules
	} else {
		response.Status = 400
		response.Message = "Error Array Size Not Connect"
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// InsertSchedule...
func InsertSchedule(w http.ResponseWriter, r *http.Request) {
	db := connect()
	defer db.Close()

	err := r.ParseForm()
	if err != nil {
		return
	}

	datetime := r.Form.Get("datetime")
	patientid, _ := strconv.Atoi(r.Form.Get("patientid"))
	medstaffid, _ := strconv.Atoi(r.Form.Get("medstaffid"))

	_, errQuery := db.Exec("INSERT INTO schedule(scheduleDateTime, patientID, medstaffID) values (?,?,?)",
		datetime,
		patientid,
		medstaffid,
	)

	var response ScheduleResponse
	if errQuery == nil {
		response.Status = 200
		response.Message = "Success"
	} else {
		response.Status = 400
		response.Message = "Insert Failed!"
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

//UpdateSchedule...
func UpdateSchedule(w http.ResponseWriter, r *http.Request) {
	db := connect()
	defer db.Close()

	err := r.ParseForm()
	if err != nil {
		return
	}

	scheduleid := r.Form.Get("id")
	datetime := r.Form.Get("datetime")
	patientid := r.Form.Get("patientid")
	medstaffid := r.Form.Get("medstaffid")

	var updating string = "update schedule set scheduleDateTime='" + datetime + "', patientID='" + patientid + "', medstaffID='" + medstaffid + "' WHERE scheduleID=" + scheduleid

	_, errQuery := db.Exec(updating)

	var response ScheduleResponse
	if errQuery == nil {
		response.Status = 200
		response.Message = "Success"
	} else {
		response.Status = 405
		response.Message = "Update Failed!"
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// DeleteSchedule...
func DeleteSchedule(w http.ResponseWriter, r *http.Request, params martini.Params) {
	db := connect()
	defer db.Close()

	err := r.ParseForm()
	if err != nil {
		return
	}

	ScheduleID := params["id"]

	_, errQuery := db.Exec("DELETE FROM schedule WHERE scheduleID=?",
		ScheduleID,
	)

	var response ScheduleResponse
	if errQuery == nil {
		response.Status = 200
		response.Message = "Success"
	} else {
		response.Status = 400
		response.Message = "Delete Failed!"
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
