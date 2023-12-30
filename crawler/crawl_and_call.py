import grpc

from . import id_page_crawler
import time
import random
import redis
import re
from . import ansChan_pb2
from . import ansChan_pb2_grpc

#   输入要爬取的网址
url = "https://www.zhihu.com/people/tuo-qia-ma-ke-zhi-guan"


def run():
    while True:
        # 生成一个50到70之间的随机数，单位是秒
        sleep_time = random.uniform(50, 70)

        # 调用函数a
        urls = id_page_crawler.crawl_main_page(url)

        new_urls = filter_urls(urls)

        data = get_ans(new_urls)

        print("grpc client sending data:", data)

        call_grpc_server(data)

        # 等待随机时间
        time.sleep(sleep_time)


def filter_urls(input_urls):
    # 创建 Redis 连接
    redis_host = '127.0.0.1'  # Redis 服务器的主机名或 IP 地址
    redis_port = 6389  # 你在 docker run 中映射的主机端口
    redis_client = redis.Redis(host=redis_host, port=redis_port, decode_responses=True)

    # 存储 URL 的键前缀
    url_key_prefix = 'url:'

    # 用于存储新 URL 的列表
    new_urls = []
    expiration_time = 604800

    for url in input_urls:
        # 构建 URL 对应的键
        url_key = f'{url_key_prefix}{url}'

        # 使用 SETNX 命令检查 URL 对应的键是否已存在
        is_set = redis_client.setnx(url_key, '1')

        if is_set:
            # 如果 URL 对应的键不存在，说明 URL 不存在于 Set 中，将其存储到 Set 中
            redis_client.expire(url_key, expiration_time)
            new_urls.append(url)

    return new_urls


def get_ans(urls):
    res = []
    for url in urls:
        html_content = id_page_crawler.get_content(url)
        if not is_target(html_content):
            continue

        author = "托卡马克之冠"

        title = extract_title(html_content)
        if title is None:
            continue

        ans = extract_ans(html_content)
        if ans is None:
            continue

        ans = wash_ans("<" + ans)

        res.append([author, title, ans])
    return res


def is_target(html_content):
    return "托卡马克之冠" in html_content


def extract_title(html_content):
    # 定义正则表达式模式
    pattern = re.compile(r'<title data-rh="true">(.+?) - 知乎</title>', re.DOTALL)

    # 在HTML内容中搜索匹配的字符串
    match = pattern.search(html_content)

    # 如果找到匹配项，则返回提取的字符串，否则返回 None
    if match:
        return match.group(1)
    else:
        return None


def extract_ans(html_content):
    # 定义正则表达式模式
    pattern = re.compile(r'<p data-first-child(.+?)</p></span></div></div></span>', re.DOTALL)

    # 在HTML内容中搜索匹配的字符串
    match = pattern.search(html_content)

    # 如果找到匹配项，则返回提取的字符串，否则返回 None
    if match:
        return match.group(1)
    else:
        return None


def wash_ans(ans):
    # 定义正则表达式模式
    pattern = re.compile(r'<.*?>')

    # 使用正则表达式替换
    result = re.sub(pattern, '\n', ans)

    return result


def call_grpc_server(data):
    with grpc.insecure_channel('127.0.0.1:1111') as channel:
        # 创建 gRPC 客户端
        stub = ansChan_pb2_grpc.AnsServiceStub(channel)

        # 创建 AnsList 消息
        ans_list = ansChan_pb2.AnsList()
        for v in data:
            ans = ans_list.arr.add()
            ans.author = v[0]
            ans.title = v[1]
            ans.content = v[2]

        # 调用 gRPC 服务端的 ProcessAnsList 方法
        try:
            response = stub.ProcessAnsList(ans_list)
            print(f"Received response: Title: {response.title}\n")
        except grpc.RpcError as e:
            print(f"Error during gRPC call: {e}")




