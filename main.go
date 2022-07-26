package main

import (
	"GoCleanMicroservice/Core"
	"GoCleanMicroservice/Delivery"
	"GoCleanMicroservice/Domain"
	"GoCleanMicroservice/GinDelivery"
	"GoCleanMicroservice/PgxRepo"
	"GoCleanMicroservice/Repo"
)

func main() {
	var pingInteractor = createPingInteractor()
	var healthRepo = createHealthRepo()
	var healthInteractor = createHealthInteractor(&healthRepo)
	var interactors = createInteractorPackage(&pingInteractor, &healthInteractor)
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

func createPingInteractor() Domain.PingInteractor {
	return &Core.PingInteractor{}
}

func createHealthInteractor(healthRepo *Repo.HealthRepo) Domain.HealthInteractor {
	return &Core.HealthInteractor{HealthRepo: healthRepo}
}

func createHealthRepo() Repo.HealthRepo {
	pool, _ := (&PgxRepo.PoolFactory{}).Create()
	return &PgxRepo.HealthRepo{Pool: pool}
}

func createServer(i *GinDelivery.InteractorPackage) Delivery.Server {
	return &GinDelivery.Server{Interactors: i}
}

func createInteractorPackage(p *Domain.PingInteractor, h *Domain.HealthInteractor) *GinDelivery.InteractorPackage {
	return &GinDelivery.InteractorPackage{Interactor: p, Health: h}
}
