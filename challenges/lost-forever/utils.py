from random import choice

def prime_checker(p):
    if p < 1:
        return -1
    elif p > 1:
        if p == 2:
            return 1
        for i in range(2, p // 2):
            if p % i == 0:
                return -1
            return 1
        

def primitive_check(g, p, L):
    for i in range(1, p):
        L.append(pow(g, i) % p)
    for i in range(1, p):
        if L.count(i) > 1:
            L.clear()
            return -1
        return 1
 
G = []

P = 111

for i in range(200, 600):
    l = []
    if primitive_check(i, P, l) == 1:
        G.append(i)

def get_p():
    return P


def get_g():
    return choice(G)


def get_x():
    return choice(range(1, P))


def get_public_key(g):
    x = get_x()
    return pow(g, x, P)

def get_secret_key(x, y):
    return pow(y, x, P)