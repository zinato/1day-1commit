### CHAPTER 32. 테스트 철학 

#### 테스트 가치 

- 테스트의 목적은 시스템에 대한 지식을 저달하는 것이고 ***각 지식은 각기 다른 가치를 지닌다.***

> 얻고자 하는 자식이 무엇인지 정확히 알아야 효과적이고 유용한 테스트를 만들 수 있다. 

#### 테스트 단언문 

- 테스트는 무언가 알기 위해 하는것이므로 반드시 ***무언가를 단언해야한다.***

> 단언문이 없으면 테스트가 아니다. 

#### 테스트 범위

> 테스트를 설계할 떄는 테스트 대상과 테스트에 속하지 않는 부분을 정확히 구분해야 한다. 

#### 테스트 가정 

> 모든 테스트는 통과, 실패, 알 수 없음, 이 세 가지 중 적어도 한 가지 결론을 도출해야 한다. 

- '알 수 없음'이라는 결과가 나왔을 때 테스트가 실패했다고 해서는 안된다. 이를 두고 실패했다고 결론을 내린다는 건 
얻은 지식이 없는데 있다고 주장하는 것이다. 

#### 테스트 설계

> 전체 테스트를 조합했을 때 원하는 지식을 빠짐없이 얻을 수 있다. 

#### E2E 테스트 

- E2E 테스트에서는 시스템으로 들어간 입력과 거기서 생산되는 결과의 양극단만 확인한다. 
- E2E 테스트의 단점은 이 테스트만으로는 시스템에 대해 알고자 하는 모든 것을 알아내기 어렵다는 것이다. 

> E2E 테스트로만 확인할 수 있는 방식으로 설계된 시스템은 코드 아키텍처에 광범위한 문제가 있다고 볼 수 있다. 

#### 통합 테스트

- 두 개 이상의 '컴포넌트'를 한 시스템에서 '조립'한 후에 어떻게 작동하는지 보는 것이 통합테스트다. 
- 여기서 말하는 컴포넌트란 코드 모듈이 될 수도 있고, 시스템이 의존하는 라이브러리 혹은 데이터를 제공하는 원격 서비스가 될 수도 있다. 
- E2E 테스트에서는 전체 시스템을 하나의 '블랙 박스'로 여기고 테스트를 진행하는 데 비해 통합 테스트에서는 컴포넌트의 분리를 중요시한다. 
- 통합 테스트만으로 시스템을 테스트  하는건 적절하지 않다. 컴포넌트의 인터랙션만으로 전체 시스템을 분석하려면 시스템 작도 방식 전반에 대해 
이해하기까지 엄청나게 많은 수의 인터랙션을 테스트해야 하기 때문이다. 
- 컴포넌트에 변화가 생기면 그 컴포넌트와 인터랙션하는 모든 컴포넌트 테스트를 업데이트해야 하므로 유지 보수 부담이 있다. 

#### 단위 테스트 

- 하나의 컴포넌트가 정상 작동하는지 확인하는 테스트가 단위 테스트다.
- 단위 테스트는 컴포넌트의 ***내부 구현***이 아니라 ***동작***을 테스트해야 한다.
- 시스템에 있는 모든 컴포넌트의 동작이 문서에 잘 정의되어 있으면 각 컴포넌트가 문서에 나온 대로 동작하는지 테스트
하기만 해도 시스템 전체의 동작을 테스트한 셈이 된다. 
- 단위 테스트는 시스템 컴포넌트가 합리적으로 잘 분리되어 있고 각 컴포넌트의 동작을 완전히 정의 할 수 있을 정도로 단순할 때 
가장 큰 효과를 낸다. 

#### 현실 

- 현실에서는 단위 테스트부터 E2E 테스트까지 무한히 많은 테스트를 해야 할 때도 있고 단위 테스트와 통합 테스트 사이, 통합 테스트와 
E2E 테스트 사이에 애매하게 걸치는 때도 있다. 현실에서는 시스템의 행동을 제대로 이해하기 위해 온갖 종류의 테스트를 
여러 단계에 걸쳐서 해야 하는게 보통이다. 

#### 결정론 

- 시스템이나 환경이 변하지 않는 한 테스트 결과는 변하지 않아야 한다. (멱등성)
- 테스트에 관한 모든게 결정론적일 필요는 없다. 시스템이 변하지 않았다면 그 테스트의 단언문은 항상 참 혹은 거짓이어야 한다. 

#### 속도 

- 테스트의 속도도 중요한 요소다. 
- 테스트 완료를 기다리느라 주의가 산만해져서 개발자가 집중력을 잃을 정도로 오랜 시간이 든다면 곤란하다.


#### 커버리지

- 실제로 커버리지를 보여주는 도구도 많이 있으나 실제 테스트를 거친 줄이 아니라 코드가 실행된 줄을 알려준다는 걸 기억하야 한다. 
- 그 코드의 동작에 관한 단언문이 없다면 실제 테스트는 거치지 않은 것이다. 

#### 결론 : 테스트의 전반적인 목표 

> 테스트의 전반적인 목표는 시스템에 관해 유효한 지식을 얻는 것이다. 

1. 어떤 명제가 참인지 확인하기 위해 작성하는 코드, 보통 명제가 거짓이면 테스트가 실패했다고 본다. 
- ex) assertEqual(the_exception.error_code, 3);
2. 소프트웨어가 목표한 동작을 정상적으로 수행하는지 확인하기 위한 테스트의 집합 