def stringCompression(chars) -> int :
    ans = 0
    i = 0

    while i < len(chars):
      letter = chars[i]
      count = 0
      while i < len(chars) and chars[i] == letter:
        count += 1
        i += 1
      chars[ans] = letter
      ans += 1
      if count > 1:
        for c in str(count):
          chars[ans] = c
          ans += 1
    print(chars)
    return ans


# def sC(chars) -> int:
#    pos = 0
#    count = 0
#    i = 0
#    while i < len(chars):
#       letter = chars[i]
#       while i < len(chars) and chars[i] == letter:
#          count +=1
#          print (chars[i],count,i)
#       i+1
#       chars[pos] = letter
#       pos +=1
#       i+=1




def main ():
    print(stringCompression(["b", "b", "b", "b"]))

if __name__ == "__main__":
    main()