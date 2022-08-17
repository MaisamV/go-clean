package main

import (
	"GoCleanMicroservice/abstract/delivery"
	"GoCleanMicroservice/abstract/domain"
	interactor2 "GoCleanMicroservice/abstract/domain/interactor"
	"GoCleanMicroservice/abstract/repo"
	"GoCleanMicroservice/app/delivery/grpc"
	"GoCleanMicroservice/app/delivery/http"
	interactor3 "GoCleanMicroservice/app/domain/app/interactor"
	"GoCleanMicroservice/app/repo/pgx"
)

func main() {
	var pingInteractor = createPingInteractor()
	var healthRepo = createHealthRepo()
	var healthInteractor = createHealthInteractor(&healthRepo)
	var interactors = createInteractorPackage(&pingInteractor, &healthInteractor)
	go func() {
		var grpcServer = createGrpcServer(interactors)
		err := grpcServer.Init()
		if err != nil {
			return
		}
		err = grpcServer.Serv()
		if err != nil {
			return
		}
	}()
	var server = createServer(interactors)
	err := server.Init()
	if err != nil {
		return
	}
	err = server.Serv()
	if err != nil {
		return
	}
}

func createPingInteractor() interactor2.PingInteractor {
	return &interactor3.PingInteractor{}
}

func createHealthInteractor(healthRepo *repo.HealthRepo) interactor2.HealthInteractor {
	return &interactor3.HealthInteractor{HealthRepo: healthRepo}
}

func createHealthRepo() repo.HealthRepo {
	pool, _ := (&pgx.PoolFactory{}).Create()
	return &pgx.HealthRepo{Pool: pool}
}

func createServer(i *domain.InteractorPackage) delivery.Server {
	return &http.Server{Interactors: i}
}

func createInteractorPackage(p *interactor2.PingInteractor, h *interactor2.HealthInteractor) *domain.InteractorPackage {
	return &domain.InteractorPackage{Interactor: p, Health: h}
}

func createGrpcServer(i *domain.InteractorPackage) delivery.Server {
	return &grpc.Server{Interactors: i}
}
