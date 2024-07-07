# Band-Protocol-Assignment-2 Superman's Chicken Rescue

## Overview
ในโจทย์ข้อที่ 2 นี้ ซูเปอร์แมนต้องปกป้องไก่จากพายุฝนโดยใช้หลังคาที่มีความยาวจำกัด ให้ข้อมูลตำแหน่งของไก่และความยาวของหลังคาที่ซูเปอร์แมนสามารถถือไปได้ จงหาจำนวนไก่สูงสุดที่ซูเปอร์แมนสามารถปกป้องได้ข้อมูลประกอบด้วยจำนวนเต็มสองตัว n และ k โดยที่ n หมายถึงจำนวนไก่และ k หมายถึงความยาวของหลังคา และบรรทัดถัดไปประกอบด้วยจำนวนเต็มชุดตัวเลขตามจำนวณของไก่(n) โดยแต่ละตัวเลขจะเป็นตำแหน่งของไก่

## Explain logic
- ทำการรับ input
- ใช้ for เพื่อหาว่าตรงไหนคือช่วงตำแหน่ง(start-end) ที่อยู่ในเงื่อนไข(k) 
- จากนั้นเช็คระยะห่างระหว่างไก่โดยเลือกระยะห่างที่ไม่เกินกระเบื้อง(k)

## Requirements
- Go 1.21.5+

## Getting Started
### Installation and Running the Project
1. Clone the project from GitHub:
   ```sh
   git clone https://github.com/PParist/band-protocol-assignment-2.git
2. Installation Go:
   ```sh
   Download and install https://go.dev/doc/install
3. Running the Project:
   ```sh
   cd assignment_2
   go run main.go / go run .
4. Test:
   ```sh
   First Input: ใส่จำนวณไก่และจำนวณกระเบื้อง
   Second Input: ใส่ระยะห่างระหว่างไก่แต่ละตัว
