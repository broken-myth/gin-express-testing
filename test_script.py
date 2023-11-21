import requests
import time


def make_get_requests(lang: str, count: int):
    for _ in range(count):
        port = "3000" if lang == "js" else "3001"
        res = requests.get(f"http://localhost:{port}/get")
        print(res.text)


def make_post_requests(lang: str, count: int):
    for _ in range(count):
        body = {"username": "johndoe", "password": "janedoe"}
        port = "3000" if lang == "js" else "3001"
        res = requests.post(f"http://localhost:{port}/post", json=body)
        print(res.text)


start_time = time.time()
make_post_requests(
    "js", 1000
)  # this function call can be changed for simulating different situations
end_time = time.time()
execution_time = (end_time - start_time) * 1000
print("Execution time:", execution_time, "milliseconds")
