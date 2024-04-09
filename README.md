# Heimdall

## Description
Do you use CLI for kubernetes? How to you switch your N kubeconfig? Don't have a solution? Ok! USE HEIMDALL!

## Architecture
A simple user can perform three action:
- List all kubeconfigs (in ~/.kube path)
- Select a kubeconfig (from list)
- Add a kubeconfig (in ~/.kube path)

```
@startuml
left to right direction
actor "User" as usr
rectangle "Heimdall Universe" {
  usecase "List Kubeconfigs" as UC1
  usecase "Select a Kubeconfig" as UC2
  usecase "Add a Kubeconfig" as UC3
}
usr --> UC1
usr --> UC2
usr --> UC3
@enduml
```

## Author
- Luca Maggio (lucamaggio1992@gmail.com)
