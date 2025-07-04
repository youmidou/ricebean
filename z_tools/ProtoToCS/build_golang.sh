#!/bin/bash
#----------------------------------------------------------------
 #生成协议 已经做到一键到位
 #added by yh @ 2023/08/03 23:35

 #protoc下载地址
 #https://github.com/protocolbuffers/protobuf/releases
 #protoc-gen-go下载地址
 #go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
 #ls $HOME/go/bin
 #复制 $HOME/go/bin 的 protoc-gen-go文件
 # use ./build_golang_mac.sh  chmod u+x *.sh
#----------------------------------------------------------

protoc_dir="./protoc"
# 指定 proto 文件所在的文件夹路径
proto_dir="./Protol"  #输入目录
output_dir="./Protocal" #指定输出目录

# 判断文件夹是否存在
if [ ! -d "$output_dir" ]; then
  # 如果文件夹不存在，就创建它
  mkdir -p "$output_dir"
  echo "创建文件夹 Folder created: $output_dir"
else
  echo "有文件夹 Folder already exists: $output_dir"
fi

echo "开始生成客户端通信协议..."

# 获取指定文件夹下的所有 proto 文件
proto_files=$(find "$proto_dir" -name "*.proto")

# 循环处理每个 proto 文件
for proto_file in $proto_files
do
  # 获取文件名（不含路径和后缀）
  base_file=$(basename "$proto_file" .proto)

  # 生成 Go 代码并指定输出目录
  $protoc_dir --proto_path="$proto_dir" --go_out="$output_dir" "$proto_file"

  # 可选：移动生成的 Go 文件到指定目录
  # mv "${output_dir}/${base_file}.pb.go" "path/to/output/directory/${base_file}.pb.go"
done

#-----------------------------------------------------------------------
echo "开始生成服务器端通信协议..."

# 指定 proto_server 文件所在的文件夹路径
proto_server_dir="./ProtolServer" #输入目录
output_server_dir="./Protocal" #指定输出目录


# 获取指定文件夹下的所有 proto 文件
proto_files=$(find "$proto_server_dir" -name "*.proto")

# 循环处理每个 proto 文件
for proto_file in $proto_files
do
  # 获取文件名（不含路径和后缀）
  base_file=$(basename "$proto_file" .proto)

  # 生成 Go 代码并指定输出目录
  $protoc_dir --proto_path="$proto_server_dir" --go_out="$output_dir" "$proto_file"

  # 可选：移动生成的 Go 文件到指定目录
  # mv "${output_dir}/${base_file}.pb.go" "path/to/output/directory/${base_file}.pb.go"
done

#-----------------------------------------------------------------------
echo "开始生成admin服务器端通信协议..."
proto_server_dir="./ProtolAdmin" #输入目录
output_server_dir="./Protocal" #指定输出目录


# 获取指定文件夹下的所有 proto 文件
proto_files=$(find "$proto_server_dir" -name "*.proto")

# 循环处理每个 proto 文件
for proto_file in $proto_files
do
  # 获取文件名（不含路径和后缀）
  base_file=$(basename "$proto_file" .proto)

  # 生成 Go 代码并指定输出目录
  $protoc_dir --proto_path="$proto_server_dir" --go_out="$output_dir" "$proto_file"

  # 可选：移动生成的 Go 文件到指定目录
  # mv "${output_dir}/${base_file}.pb.go" "path/to/output/directory/${base_file}.pb.go"
done

echo "生成协议文件成功！！！"

#----------todo---生成协议给前端用-------------------------------------
echo "cp 协议 to unity客户端 ..."
protol_dir="/Users/yh/Documents/github/crazy-elimination/Assets/Hotfix/Net/Protol"
protocal_dir="/Users/yh/Documents/github/crazy-elimination/Assets/Hotfix/Net/Protocal"
echo "cp config..."
rm -rf "$protol_dir"
mkdir "$protol_dir"
cp -rf Protol/ "$protol_dir/"

echo "生成客户端协议..."
./build_csharp.sh
cp -rf ProtocalCS/Protol/ "$protocal_dir/"
echo "生成客户端协议...完成"
#----------todo---生成协议给web_admin前端用-------------------------------------

echo "cp admin协议 to webserver ..."
protol_admin_dir="/Users/yh/Documents/github/phoenix-tudou-web/z_Tools/ProtoToCS/ProtolAdmin"
echo "cp config..."
rm -rf "$protol_admin_dir"
mkdir "$protol_admin_dir"

cp -rf ProtolAdmin/ "$protol_admin_dir/"

