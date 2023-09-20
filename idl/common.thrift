namespace go common

enum Role {
  SuperAdmin,
  Admin,
  User,
  Guest,
}

exception RequestException{
    required i64 code;
    string msg;
}

struct EmptyStruct{}