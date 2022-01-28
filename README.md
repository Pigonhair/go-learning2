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

