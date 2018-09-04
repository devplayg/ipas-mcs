package objs

type IpasTrackingFilter struct {
	IpasFilter
	EquipIdWithOrg string  `form:"equip_id_with_org"`
}