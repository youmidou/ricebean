#!/bin/bash


# 指定 proto 文件所在的文件夹路径
proto_dir="./Protol"
# 指定输出目录
output_dir="./ProtocalCS"
protoc_dir="./protoc"

# 判断文件夹是否存在
if [ ! -d "$output_dir" ]; then
  # 如果文件夹不存在，就创建它
  mkdir -p "$output_dir"
  echo "创建文件夹 Folder created: $output_dir"
else
  echo "有文件夹 Folder already exists: $output_dir"
fi

echo 读取版本
$protoc_dir --version
#读取版本

echo "开始生成客户端通信协议..."

# 获取指定文件夹下的所有 proto 文件
proto_files=$(find "$proto_dir" -name "*.proto")

# 循环处理每个 proto 文件
for proto_file in $proto_files
do
  # 获取文件名（不含路径和后缀）
  base_file=$(basename "$proto_file" .proto)
  #echo "base_file->" $base_file "/" $proto_file
  # 生成 C# 代码并指定输出目录
  $protoc_dir  +names=original --proto_path="$proto_dir" --csharp_out="$output_dir" "$proto_file"
  #$protoc_dir +names=original --proto_path="$proto_dir" --csharp_out="$output_dir" "$proto_file"
  # 可选：移动生成的 Go 文件到指定目录
  # mv "${output_dir}/${base_file}.pb.go" "path/to/output/directory/${base_file}.pb.go"
done


echo "生成协议文件成功！！！"

#protol_dir="/Users/yh/Documents/github/crazy-elimination/Assets/Hotfix/Net/Protol"
#echo "cp config..."
#rm -rf "$protol_dir"
#mkdir "$protol_dir"

#cp -rf Protol/ "$protol_dir/"