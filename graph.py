import pandas as pd
import matplotlib.pyplot as plt
import glob
import re
import os

BASE_DIR = os.path.dirname(os.path.abspath(__file__))
DATA_DIR = os.path.abspath(os.path.join(BASE_DIR, "..", "data")) #sesuaikan dengan path csv-nya

files = glob.glob(os.path.join(DATA_DIR, "result_n*.csv"))

def extract_n(filename):
    match = re.search(r"result_n(\d+)\.csv", filename)
    return int(match.group(1))

files = sorted(files, key=extract_n)

print("CSV digunakan:")
for f in files:
    print(" -", f)

N = []
time_iter = []
time_recur = []

for f in files:
    df = pd.read_csv(f)

    N.append(int(df["N"][0]))
    time_iter.append(float(df["Time_Iter"][0]))
    time_recur.append(float(df["Time_Recur"][0]))

print("Urutan N:", N)

plt.figure()
plt.plot(N, time_iter, marker='o', label='Iteratif')
plt.plot(N, time_recur, marker='o', label='Rekursif')
plt.xlabel("Ukuran Data (n)")
plt.ylabel("Waktu Eksekusi (detik)")
plt.title("Perbandingan Pencarian Nilai Minimum")
plt.legend()
plt.grid(True)
plt.show()

plt.figure()
plt.plot(N, time_iter, marker='o', label='Iteratif')
plt.plot(N, time_recur, marker='o', label='Rekursif')
plt.xlabel("Ukuran Data (n)")
plt.ylabel("Waktu Eksekusi (detik)")
plt.title("Perbandingan Pencarian Nilai Maksimum")
plt.legend()
plt.grid(True)
plt.show()

plt.figure()
plt.plot(N, time_iter, marker='o', label='Iteratif (Min+Max)')
plt.plot(N, time_recur, marker='o', label='Rekursif (Min+Max)')
plt.xlabel("Ukuran Data (n)")
plt.ylabel("Waktu Eksekusi (detik)")
plt.title("Kinerja Keseluruhan Algoritma")
plt.legend()
plt.grid(True)
plt.show()
