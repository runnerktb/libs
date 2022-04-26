package universe

import (
	"strings"
)

var (
	noun = []string{
		"_anomaly",
		"_apps",
		"_checkpoint",
		"_customer",
		"_device",
		"_device_activation",
		"_fuel",
		"_functions",
		"_geo_group",
		"_geocode",
		"_geofence",
		"_job_management",
		"_notif_abnormal",
		"_notif_anomaly",
		"_notif_disclaimer",
		"_notif_event",
		"_notif_geofence",
		"_notif_maintenance",
		"_notif_poi",
		"_notif_promo",
		"_poi",
		"_realms",
		"_realtime",
		"_refuel",
		"_report_customer_activation",
		"_report_event",
		"_report_geofence",
		"_report_idling",
		"_report_overspeed",
		"_report_site",
		"_report_stop",
		"_report_towing",
		"_report_trip",
		"_report_vehicle_activation",
		"_report_working_hour",
		"_services",
		"_track_replay",
		"_trip",
		"_user",
		"_user_access",
		"_user_org",
		"_user_type",
		"_vehicle",
		"_vehicle_category",
		"_vehicle_group",
		"_vehicle_history",
		"_vehicle_model",
		"_vehicle_starter",
		"_vehicle_status",
		"_vehicle_type",
		"_sparepart",
		"_group_sparepart",
		"_subgroup_sparepart",
	}
)

type Access struct {
	ID     string
	Action string
}

func ParseModule(module string) (res Access) {
	if id, ok := GetModule(module); ok {
		res.ID = id
	}
	for _, v := range noun {
		r := strings.SplitN(module, v, 2)
		if len(r) == 2 && len(r[0]) > 0 {
			res.Action = r[0]
			return
		}
	}
	return
}
