#! /usr/bin/env bash

module="github.com/Ocyss/Qblog"

idls=("article"
      "auth"
      "comment"
      "notification"
      "search"
      "statistics"
      "storage"
      "user"
      )

function gen() {
  idl=$1
  kitex -module $module idl/$idl.thrift
  (
    mkdir -p cmd/$idl ; cd $_
    kitex -module $module -use $module/kitex_gen -template-dir ../../template ../../idl/$idl.thrift
  )
#  ( cd cmd/api && hz update --idl=../../idl/$idl.thrift --customize_package=template/package.yaml )
}

for idl in "${idls[@]}"; do
  gen "$idl"
done