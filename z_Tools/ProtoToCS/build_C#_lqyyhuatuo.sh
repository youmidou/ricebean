#!/bin/bash

# 指定 proto 文件所在的文件夹路径
proto_dir="./Protol"

# 指定输出目录
output_dir="/Users/yh/Documents/github/crazy-elimination/Assets/Hotfix/Net/Protocal"
protol_dir="/Users/yh/Documents/github/crazy-elimination/Assets/Hotfix/Net/Protol"

protoc_dir="./protoc"

rm -rf "$output_dir"
mkdir "$output_dir"

# 获取指定文件夹下的所有 proto 文件
proto_files=$(find "$proto_dir" -name "*.proto")

# 循环处理每个 proto 文件
for proto_file in $proto_files
do
  # 获取文件名（不含路径和后缀）
  base_file=$(basename "$proto_file" .proto)

  # 生成 C# 代码并指定输出目录
  #$protoc_dir --proto_path="$proto_dir" --csharp_out="$output_dir" "$proto_file"
  #$protoc_dir --csharp_out="$output_dir" "$proto_file"
  echo "生成$output_dir/$base_file.cs"
  # 可选：移动生成的 C# 文件到指定目录
  # mv "${output_dir}/${base_file}.cs" "path/to/output/directory/${base_file}.cs"
done

echo "cp config..."
rm -rf "$protol_dir"
mkdir "$protol_dir"

cp -rf Protol/ "$protol_dir/"