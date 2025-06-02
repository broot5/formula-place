def validate_empty_str_to_none(v: str | None):
    return None if v == "" else v
