def moveZero(nums):
    for i in nums:
        if i == 0:
            nums.remove(i)
            nums.append(0)
    return nums



def moveZeroNew(nums):
    j=0
    n=len(nums)

    for i in range (n):
        if nums[i] != 0:
            nums[j]=nums[i]
            j += 1
        # print(i)
    
    for j in range(j,n):
        nums[j]=0
        # print(j)
    return nums


def main():
    print(moveZeroNew([0,1,0,3,12]))
if __name__ == "__main__":
    main()