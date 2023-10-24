from pettingzoo.classic import rps_v2
import json

env = rps_v2.env()
env.reset(seed=42)
acc_rewards = env.rewards.copy()

itt = env.agent_iter()
for agent in env.agent_iter():
    observation, reward, termination, truncation, info = env.last()
    print(json.dumps({'agent':agent, 'observation':observation.tolist()}))
    if termination or truncation:
        action = None
    else:
        # this is where you would insert your policy
        action = env.action_space(agent).sample()

    env.step(action)
    for agent, reward in env.rewards.items():
        acc_rewards[agent]+=reward
env.close()
print(json.dumps(acc_rewards))
