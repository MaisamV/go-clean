package main

import (
	"GoCleanMicroservice/abstract/delivery"
	"GoCleanMicroservice/abstract/domain"
	"GoCleanMicroservice/abstract/domain/interactor"
	"GoCleanMicroservice/abstract/repo"
	"GoCleanMicroservice/app/delivery/grpc"
	"GoCleanMicroservice/app/delivery/http"
	app "GoCleanMicroservice/app/domain/app/interactor"
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

func createPingInteractor() interactor.PingInteractor {
	return &app.PingInteractor{}
}

func createHealthInteractor(healthRepo *repo.HealthRepo) interactor.HealthInteractor {
	return &app.HealthInteractor{HealthRepo: healthRepo}
}

func createHealthRepo() repo.HealthRepo {
	pool, _ := (&pgx.PoolFactory{}).Create()
	if pool == nil {
		panic(any("Couldn't create database pool. Check DATABASE_URL environment variable, connection string and access to server."))
	}
	return &pgx.HealthRepo{Pool: pool}
}

func createServer(i *domain.InteractorPackage) delivery.Server {
	return &http.Server{Interactors: i}
}

func createInteractorPackage(p *interactor.PingInteractor, h *interactor.HealthInteractor) *domain.InteractorPackage {
	return &domain.InteractorPackage{Ping: p, Health: h}
}

func createGrpcServer(i *domain.InteractorPackage) delivery.Server {
	return &grpc.Server{Interactors: i}
}
