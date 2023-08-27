# def knapsack(W, wt, val):
#   print(W)
#   print(wt)
#   n = len(wt)
#   K = [[0 for x in range(W + 1)] for x in range(n + 1)]
#   print(K)
#   for i in range(n + 1):
#     for w in range(W + 1):
#       print("i={i}, w={w}, wt[i-1]={wt}".format(i = i, w = w, wt=wt[i-1]))
#       if i == 0 or w == 0:
#         K[i][w] = 0
#       elif wt[i-1] <= w:
#         K[i][w] = max(val[i-1] + K[i-1][w - wt[i-1]], K[i-1][w])
#       else:
#         K[i][w] = K[i-1][w]
#   print(K)
#   return K[n][W]

# res = knapsack(5, [2,3,4], [4,3,2])
# print("res=", res)

def knapsack(W, wt, val):
    n = len(wt)
    K = [[0 for x in range(W + 1)] for x in range(n + 1)]
    A = [[[] for x in range(W + 1)] for x in range(n + 1)]
    print(K)
    print(A)
    for i in range(n + 1):
        for w in range(W + 1):
            if i == 0 or w == 0:
                K[i][w] = 0
            elif wt[i - 1] <= w:
                c = val[i - 1] + K[i - 1][w - wt[i - 1]]
                p = K[i - 1][w]
                if c > p:
                    K[i][w] = c
                    A[i][w] = A[i - 1][w - wt[i - 1]][:]
                    A[i][w].append(i - 1)
                else:
                    K[i][w] = p
                    A[i][w] = A[i - 1][w]
            else:
                K[i][w] = K[i - 1][w]
                A[i][w] = A[i - 1][w]
    print(K)
    return K[n][W], A[n][W]


(res, indexes) = knapsack(5, [3, 2, 2], [3, 2, 4])
print("res=", res)
print("indexes=", indexes)
