#!/bin/bash

# 获取当前脚本所在目录的绝对路径
current_dir=$(cd $(dirname $0) && pwd)

# 输出文件名
output_file="${current_dir}/all_go_files.txt"

# 删除旧的输出文件
rm -f "${output_file}"

# 遍历当前目录的父目录下的所有目录，输出其中的go代码文件
find "${current_dir}/../" -type f -name "*.go" | while read file; do
  # 获取文件相对于父目录的路径
  relative_path="${file#${current_dir}/../}"

  # 输出文件路径
  echo "${relative_path}" >> "${output_file}"

  # 删除文件中的空行并输出到输出文件中
  sed '/^\s*$/d' "${file}" >> "${output_file}"
done
