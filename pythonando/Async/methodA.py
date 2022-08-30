import asyncio

async def methodA():
    await asyncio.sleep(4)
    print("methodA")

async def methodB():
    await asyncio.sleep(1)
    print("methodB")
    task2 = asyncio.create_task(methodA())
    await task2

async def methodC():
    await asyncio.sleep(3)
    for _ in range(3):
        print("methodC")
        await asyncio.sleep(1)
        
async def teste():
    task = asyncio.create_task(methodB())
    task1 = asyncio.create_task(methodC())
    await task
    await task1
    

asyncio.run(teste())