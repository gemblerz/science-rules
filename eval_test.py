import numpy as np

measurements = [
    {"name": "env.temperature", "value": 23.1, "meta": {"sensor": "bme680"}},
    {"name": "env.temperature", "value": 41.2, "meta": {"sensor": "bme280"}},
    {"name": "sys.uptime", "value": 3700, "meta": {}},
]


# !!! main function of interest compact notation for selecting measurements
def v(name, **meta):
    return np.array([m["value"] for m in measurements if m["name"] == name and matchmeta(meta, m["meta"])])


def matchmeta(pattern, meta):
    return all(k in meta and meta[k] == v for k, v in pattern.items())


def avg(vs):
    return np.mean(vs)


# select all values
print(v('env.temperature'))

# select values from bme680
print(v('env.temperature', sensor='bme680'))

# select values from bme280
print(v('env.temperature', sensor='bme280'))

# since the returned value is a numpy array, you can also broadcast operations
print(v('env.temperature', sensor='bme280') > 40.0)

# you can eval as usual in this environment
print("eval 1", eval("""
v('env.temperature', sensor='bme280') > 40.0
"""))

print("avg", eval("""
avg(v('env.temperature')) > 40.0
"""))

# finally, you can post evaluate a rule to do an all / any check
print("all?", eval("""
all(v('env.temperature') > 40.0)
"""))

print("any?", eval("""
any(v('env.temperature') > 20.0)
"""))

eval("1 + 2")

import time

eval("time.time()")
# TODO see how to fit time windows