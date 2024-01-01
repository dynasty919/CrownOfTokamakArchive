# 使用 Python 3.6.9 镜像作为基础镜像
FROM python:3.6.9

# 设置工作目录
WORKDIR /app

# 添加当前目录到 Python 路径中
ENV PYTHONPATH /app/crawler

# 复制项目文件到容器中
COPY ./crawler /app/crawler

# 安装 Python 依赖
RUN pip install --no-cache-dir -r /app/crawler/requirements.txt

# 暴露所需的端口（如果有需要）
# EXPOSE 5000

# 定义容器启动命令
CMD ["python", "-m", "crawler.entrance"]