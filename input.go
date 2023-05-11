package main

// input column data here
type Column struct {
	Column1  Data `json:"id" dataType:"int" tribeOwner:"3" isPII:"false" typePII:""`
	Column2  Data `json:"quest_user_id" dataType:"int" tribeOwner:"3" isPII:"false" typePII:""`
	Column3  Data `json:"benefit_type" dataType:"int" tribeOwner:"3" isPII:"false" typePII:""`
	Column4  Data `json:"benefit_unique_code" dataType:"string" tribeOwner:"3" isPII:"false" typePII:""`
	Column5  Data `json:"status" dataType:"int" tribeOwner:"3" isPII:"false" typePII:""`
	Column6  Data `json:"metadata" dataType:"json" tribeOwner:"3" isPII:"false" typePII:""`
	Column7  Data `json:"created_at" dataType:"int" tribeOwner:"3" isPII:"false" typePII:""`
	Column8  Data `json:"created_by" dataType:"int" tribeOwner:"3" isPII:"true" typePII:"C3"`
	Column9  Data `json:"updated_at" dataType:"int" tribeOwner:"3" isPII:"false" typePII:""`
	Column10 Data `json:"updated_by" dataType:"int" tribeOwner:"3" isPII:"true" typePII:"C3"`
}

// input table data here
type Table struct {
	TableName Column `json:"quest.quest_trx_quest_user_benefit"`
}
