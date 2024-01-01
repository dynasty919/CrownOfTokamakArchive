from datetime import datetime
import requests
import re


# 获取网页body里的内容
def get_content(url, data=None):
    # 设置Http请求头
    headers = {
        'Accept': 'text/template,application/xhtml+xml,application/xml;q=0.9,image/webp,*/*;q=0.8',
        'Accept-Encoding': 'gzip, deflate, sdch',
        'Accept-Language': 'zh-CN,zh;q=0.8',
        'Connection': 'keep-alive',
        'User-Agent': 'Mozilla/5.0 (Windows NT 6.3; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/43.0.235'
    }

    proxies = {
        'http': '192.168.56.1:7890',
        'https': '192.168.56.1:7890',
    }

    # 获取当前时间
    current_time = datetime.now().strftime("%Y-%m-%d %H:%M:%S")
    print(f"try to crawl {url}, at{current_time}")

    req = requests.get(url, headers=headers)
    req.encoding = 'utf-8'
    return req.text


def extract_contents(html_content):
    pattern = r'"type":"answer".*?"type":"question"'
    matches = re.findall(pattern, html_content)
    return matches


def extract_id(kv_pairs):
    result_list = []
    for item in kv_pairs:
        pair = []
        arr = item.split(',')

        for s in arr:
            if len(s) >= 6 and s[:6] == '"url":':
                match = re.search(r'(\d+)"$', s)
                if match:
                    pair.append(match.group(1))

        result_list.append(pair)

    return result_list


def get_answers_urls(id_pairs):
    result = []
    for pair in id_pairs:
        result.append("https://www.zhihu.com/question/" + pair[1] +
                      "/answer/" + pair[0])
    return result


def crawl_main_page(url):
    html_content = get_content(url)
    kv_pairs = extract_contents(html_content)
    id_pairs = extract_id(kv_pairs)
    answers_urls = get_answers_urls(id_pairs)
    return answers_urls
