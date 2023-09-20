namespace go common

enum Role {
  SuperAdmin,
  Admin,
  User,
  Guest,
}

exception RequestException{
    required i32 code;
    string msg;
}

struct EmptyStruct{}