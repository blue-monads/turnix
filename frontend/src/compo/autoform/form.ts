export interface Schema {
  name: string;
  fields: Field[];
  required_fields: string[];
}

export interface Field {
  name: string;
  key_name: string;
  ftype:
    | "TEXT"
    | "SELECT"
    | "TEXT_SLUG"
    | "MULTI_TEXT"
    | "LONG_TEXT"
    | "TEXT_POLICY"
    | "BOOL"
    | "INT"
    | "KEY_VALUE_TEXT"
    | "USER_GROUP"
    | "USER"
    | "BPRINT"
    | "PLUG"
    | "AGENT"
    | "HANDLER"
    | "RESOURCE";

  linked_fields?: string[];
  disabled?: boolean;
  options?: string[];
  slug_gen?: () => string;
}
