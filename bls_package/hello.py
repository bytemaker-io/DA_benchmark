from Cryptodome.Math.Numbers import Integer
from Cryptodome.Math.Polynomial import Polynomial
from Cryptodome.Math.Polynomial.ECCurve import CurveFp
from Cryptodome.Math.Polynomial.ECPoint import Point
from Cryptodome.Util.number import long_to_bytes

# 定义椭圆曲线参数
p = 0xFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFEFFFFFC2F  # 素数 p
n = 0xFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFEBAAEDCE6AF48A03BBFD25E8CD0364141  # 群阶
a = 0x0000000000000000000000000000000000000000000000000000000000000000  # 系数 a
b = 0x0000000000000000000000000000000000000000000000000000000000000007  # 系数 b
G_x = 0x79BE667EF9DCBBAC55A06295CE870B07029BFCDB2DCE28D959F2815B16F81798  # 基点 G 的 x 坐标
G_y = 0x483ADA7726A3C4655DA4FBFC0E1108A8FD17B448A68554199C47D08FFB10D4B8  # 基点 G 的 y 坐标

# 创建椭圆曲线对象
curve = CurveFp(p, a, b)
G = Point(curve, Integer(G_x), Integer(G_y))

# 生成随机的标量值和点
d1 = Integer.random_range(1, n)
d2 = Integer.random_range(1, n)
h1 = Point(curve, Integer.random_range(1, n), Integer.random_range(1, n))
h2 = Point(curve, Integer.random_range(1, n), Integer.random_range(1, n))

# 计算双线性对左侧
left_side = curve.pair(G * (d1 + d2), h1 + h2)

# 计算双线性对右侧
right_side = curve.pair(G, d1 * h1 + d2 * h2)

# 验证式子是否成立
if left_side == right_side:
    print("式子成立")
else:
    print("式子不成立")
