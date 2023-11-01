import sys
import json

def agent(state: dict) -> dict:
    return {"other": 5}

if __name__=='__main__':
    res = agent(json.loads(sys.argv[1]))
    print(json.dumps(res))
