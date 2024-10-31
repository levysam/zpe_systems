table "roles" {
  schema = schema.zpe
  column "p_type" {
    null    = false
    type    = varchar(32)
    default = ""
  }
  column "v0" {
    null    = false
    type    = varchar(255)
    default = ""
  }
  column "v1" {
    null    = false
    type    = varchar(255)
    default = ""
  }
  column "v2" {
    null    = false
    type    = varchar(255)
    default = ""
  }
  column "v3" {
    null    = false
    type    = varchar(255)
    default = ""
  }
  column "v4" {
    null    = false
    type    = varchar(255)
    default = ""
  }
  column "v5" {
    null    = false
    type    = varchar(255)
    default = ""
  }
  index "idx_roles" {
    columns = [column.p_type, column.v0, column.v1]
  }
}
table "users" {
  schema = schema.zpe
  column "ID" {
    null = false
    type = char(26)
  }
  column "Name" {
    null = false
    type = char(255)
  }
  column "Password" {
     null = false
     type = char(255)
  }
  column "Email" {
    null = false
    type = char(255)
  }
  primary_key {
    columns = [column.ID]
  }
  index "idx_users" {
      columns = [column.Email]
      unique = true
    }
}
schema "zpe" {
  charset = "utf8mb4"
  collate = "utf8mb4_0900_ai_ci"
}
