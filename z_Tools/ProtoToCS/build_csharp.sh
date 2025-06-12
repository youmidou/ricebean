#!/bin/bash
#build_csharp
#added by yh @ 2024-03-13 17:22
current_path=$(pwd)
echo "当前路径是：$current_path"

#protoc_dir="./csharp/net6.0/protogen"
protoc_dir="./csharp_net6.0/protogen"
# 指定 proto 文件所在的文件夹路径
proto_dir="./Protol"
# 指定输出目录
output_dir="./ProtocalCS"
#
target_dir="./csharp_net6.0"

rm -rf "$output_dir"
mkdir "$output_dir"

echo "开始生成客户端通信协议..."
#复制Protol->csharp_net6.0/
cp -rf Protol/ "$target_dir/Protol"

# 获取指定文件夹下的所有 proto 文件
proto_files=$(find "$proto_dir" -name "*.proto")

# 循环处理每个 proto 文件
for proto_file in $proto_files
do
  # 获取文件名（不含路径和后缀）
  base_file=$(basename "$proto_file" .proto)
  # 上一级目录中的文件
  #echo "原始路径= $proto_dir"
  # 生成 csharp 代码并指定输出目录
  $protoc_dir +names=original --proto_path="$proto_dir" --csharp_out="$output_dir" "$proto_file"
  #$protoc_dir +names=original --proto_path="$proto_dir" --csharp_out="$output_dir" "$proto_file"

done

echo "Protocol generation completed."

rm -rf "$target_dir/Protol"