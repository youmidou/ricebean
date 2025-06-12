#!/bin/bash

# 指定 proto 文件所在的文件夹路径
proto_dir="path/to/your/proto/directory"

# 指定输出目录
output_dir="path/to/output/directory"

# 获取指定文件夹下的所有 proto 文件
proto_files=$(find "$proto_dir" -name "*.proto")

# 循环处理每个 proto 文件
for proto_file in $proto_files
do
  # 获取文件名（不含路径和后缀）
  base_file=$(basename "$proto_file" .proto)

  # 生成 Go 代码并指定输出目录
  protoc --go_out="$output_dir" "$proto_file"

  # 可选：移动生成的 Go 文件到指定目录
  # mv "${output_dir}/${base_file}.pb.go" "path/to/output/directory/${base_file}.pb.go"
done
