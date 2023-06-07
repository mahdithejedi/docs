def all_states(n):
    return pow(2, n)

def zero_near_state(n):
    return n - 2

num_count = int(input())
print(all_states(num_count) - zero_near_state(num_count))
