import asyncio
import aiohttp
import random

base_URL = "http://43.139.176.247:30010/api/"
subfix = [
    'ping',
    'info',
    'sites',
    'login'
]

async def fake_visit(times, duration):
    async with aiohttp.ClientSession() as session:
        for _ in range(times):
            await session.get(base_URL + random.choice(subfix))
            await asyncio.sleep(duration)

async def main():
    tasks = []
    n = random.randint(1, 20)
    duration = random.random()
    for i in range(n):
        tasks.append(fake_visit(n, duration))
    await asyncio.gather(*tasks)


if __name__ == '__main__':
    print('模拟随机访问开始.....')
    asyncio.run(main())
    print('模拟随机访问结束.....')
