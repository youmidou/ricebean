#!/bin/bash
#build_csharp
#added by yh @ 2024-03-06 12:11

#protoc_dir="./csharp/net6.0/protogen"
protoc_dir="./protogen"
# 指定 proto 文件所在的文件夹路径
proto_dir="./Protol"
# 指定输出目录
output_dir="./ProtocalCS"

rm -rf "$output_dir"
mkdir "$output_dir"

echo "开始生成客户端通信协议..."

# 获取指定文件夹下的所有 proto 文件
proto_files=$(find "$proto_dir" -name "*.proto")

# 循环处理每个 proto 文件
for proto_file in $proto_files
do
  # 获取文件名（不含路径和后缀）
  base_file=$(basename "$proto_file" .proto)
  #echo "= $proto_file"
  # 生成 csharp 代码并指定输出目录
  $protoc_dir +names=original --proto_path="$proto_dir" --csharp_out="$output_dir" "$proto_file"
  #$protoc_dir +names=original --proto_path="$proto_dir" --csharp_out="$output_dir" "$proto_file"

done

echo "Protocol generation completed."

