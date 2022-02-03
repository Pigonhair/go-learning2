# Golang 개념&공부    
## 2022-01-28~
## Go 장점 및 특징   
* 간결한 문법 및 단순함   
* 병행 프로그래밍 지원
* 정적 타입 및 동적 실행
* 간편한 협업 지원
* 컴파일 및 실행속도 빠름
* 제네릭 및 예외 처리 미지원
* 컨벤션 통일

## 변수 및 상수 선언법   
* 변수 선언 및 사용법: var
* 상수 선언 및 사용법: const

## Go에서 package개념
* 최초실행 package는 main이 되야한다(자바의 public class main과 같은개념) -> main 패키지가 아니면 실행 안됨
* 그외에 package들은 main package에서 import해서 사용

## 변수 선언(일반적인 선언 방법)
* var j string = "Hi ! Golang!"

## 변수 선언(짧은 선언)
* 함수 안에서만 사용(전역X), 선언 후 할당 예외 발생
* 주로 제한된 범위의 함수내에서 사용 할 경우 코드 가독성을 높일 수 있다
* shortVar1 := "Test"
* shortVar2 := false

## 열거형(Enumeration)
* 상수를 사용하는 일정한 규칙에 따라 숫자를 계산 및 증가시키는 묶음
*
   ``` go
       const ( 
          A = iota
          B
          C
        )
   ```
   * 하면 A는 0, B는 1, C는 2 값이 저장됨
   * iota는 0부터 순차적으로 증가하게 해주는 예약어
   * 여기서 A대신에 _를 입력하면 B와 C에만 값이 저장된다(_는 생략의 의미)
## 제어문 및 반복문(IF, SWITCH, FOR)
### 제어문(조건문)
* IF 문 : 반드시 Boolean 검사 -> 1, 0 (사용불가 : 자동 형 변환 불가)
* 소괄호 사용 x
### SWITCH CASE에서 fallthrough
* 각 case가 끝날때 fallthrough를 넣어주면 조건에 맞지않아도 그 다음 case문이 실행됨
### 반복문 - For
* Go에서 유일하게 제공되는 반복문
* 무한 루프(자바의 while문)로 하고 싶으면 for문에 조건식을 안쓰면 된다
   * 빠져나올때는 자바와 마찬가지로 break
* Range용법    
   ```go
         loc := []string{"Seoul", "Busan", "Incheon"}
	      for index, name := range loc {
		      fmt.Println("ex3 : ", index, name)
	       }
   ```
   * index와 해당 값 순서대로 출력됨(index는 _같은걸로 생략해서 값만 보이게 출력 할 수 있음)
* j := i++ 와 같은 Go에서 후치연산은 반환 값이 없기 때문에 불가능한 코드
