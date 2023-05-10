# Algorithmic complexity?
Algorithm complexity is an estimate of the growth in algorithmic execution time and storage space required, usually measured in Big O notation.
The main factors to evaluate an algorithm include: execution time, storage space, programmer's effort.

Độ phức tạp của thuật toán là đánh giá về sự tăng trưởng của thời gian thực hiện thuật toán và không gian lưu trữ cần sử dụng, thường được đo bằng Big O notation.
Các yếu tố chính để đánh giá một thuật toán bao gồm: thời gian thực thi, không gian lưu trữ, công sức thực hiện của lập trình viên.

<div align="center">
    <img src="images/big-o-notation"></img>
</div>

## Big O notation?
Big O notation is a way of expressing the upper bound of the growth rate of an algorithm's execution time or storage space as the input data size increases infinitely. Or simply the worst-case scenario, consuming the most time and resources.

Big O notation là một cách thể hiện giới hạn trên của tỷ lệ tăng trưởng của thời gian thực thi hoặc không gian lưu trữ của thuật toán khi kích thước dữ liệu đầu vào tăng lên vô hạn. Hay đơn giản là trường hợp xấu nhất có thể xảy ra, gây tốn thời gian và tài nguyên nhất.

Commonly encountered complexities in the algorithm:
Time complexity is a measure of the execution time of an algorithm as a function of its input size. It describes how the running time of an algorithm increases as the input size increases.
Space complexity is a measure of the amount of memory used by an algorithm during execution, as a function of its input size. It describes how the amount of memory used by an algorithm increases as the input size increases.

Các độ phưc tạp thường gặp trong thuật toán:
- Time complexity(độ phức tạp thời gian) là một đo lường về thời gian thực hiện một thuật toán dưới dạng một hàm số của kích thước đầu vào của nó. Nó mô tả cách thời gian chạy của một thuật toán tăng lên khi kích thước đầu vào tăng lên.
- Space complexity(độ phức tạp không gian) là một đo lường về khối lượng bộ nhớ được sử dụng bởi một thuật toán trong quá trình thực thi, dưới dạng một hàm số của kích thước đầu vào của nó. Nó mô tả cách lượng bộ nhớ được sử dụng bởi một thuật toán tăng lên khi kích thước đầu vào tăng lên.

Some common complexity symbols:
- O(1): constant complexity - execution time does not depend on input data size
- O(logN): logarithmic complexity - execution time increases with the logarithm of the input data size
- O(N): linear complexity - linear execution time with input data size
- O(NlogN): logarithmic linear complexity - execution time increases as a function of n logarithms of input data size
- O(N^2): complexity order 2 - linear execution time with squared input data method
- O(N^3): 3rd order complexity - linear execution time with 3rd order input data size
- O(B^N) B > 1: exponential complexity - execution time increases with function B^N input data size
- O(N!): huge complexity - execution time increases with n! input data size

Một số ký hiệu độ phức tạp thường gặp: 
- O(1): độ phức tạp hằng số - thời gian thực thi không phụ thuộc vào kích thước dữ liệu đầu vào
- O(logN): độ phức tạp logarithmic - thời gian thực thi tăng theo hàm logarit của kích thức dữ liệu đàu vào
- O(N): độ phức tạp tuyến tính - thời gian thực thi tuyến tính với kích thước dữ liệu đầu vào 
- O(NlogN): độ phức tạp tuyến tính logarithmic -  thời gian thực thi tăng theo hàm n logarit của kích thước dữ đầu vào
- O(N^2): độ phức tạp bậc 2 - thời gian thực thi tuyến tính với bình phương kích thức dữ liệu đầu vào
- O(N^3): độ phức tạp bậc 3 - thời gian thực thi tuyến tính với bậc 3 kích thức dữ liệu đầu vào
- O(B^N) B > 1: độ phức tạp mũ - thời gian thực thi tăng theo hàm B^N kích thức dữ liệu đầu vào
- O(N!): độ phức tạp cực lớn - thời gian thực thi tăng theo n! kích thước dữ liệu đầu vào

