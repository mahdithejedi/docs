package main

import (
	"fmt"
	"github.com/fatih/color"
	"log"
)

func main() {
	color.Set(color.Red)
	defer color.Uset()
	defer func() { fmt.Println("Goodby") }()
	log.Fatal("You can not use OOP pattern in struct langauge")
}

//
//import "fmt"
//
//// ---- Device
//
//type device interface {
//	Enable()
//	Disable()
//	IsEnable() bool
//	SetVolume(volume uint8)
//	GetVolume() uint8
//	IncreaseVolume()
//	DecreaseVolume()
//}
//
//type SystemType string
//
//const (
//	windows SystemType = "Windows"
//	linux   SystemType = "Linux"
//	mac     SystemType = "Mac"
//)
//
//type Devise struct {
//	on         bool
//	volume     uint8
//	DeviceType SystemType
//	device
//}
//
//type Windows struct {
//	Devise
//}
//
//func (w *Windows) Enable() {
//	w.Devise.on = true
//}
//
//func (w *Windows) Disable() {
//	w.Devise.on = false
//}
//func (w *Windows) IsEnable() bool {
//	return w.Devise.on
//}
//
//func (w *Windows) SetVolume(volume uint8) {
//	if volume > 100 {
//		w.Devise.volume = 100
//	}
//	w.Devise.volume = volume
//}
//
//func (w *Windows) GetVolume(volume uint8) uint8 {
//	return w.Devise.volume
//}
//
//func (w *Windows) IncreaseVolume() {
//	if w.Devise.volume >= 100 {
//		return
//	}
//	w.Devise.volume++
//}
//
//func (w *Windows) DecreaseVolume() {
//	if w.Devise.volume <= 0 {
//		return
//	}
//	w.Devise.volume--
//}
//
//type remote interface {
//	MaxVolume()
//	UnVolume()
//	TurnOff()
//}
//
//type Remote struct {
//	DeviceInstance struct {
//		Devise
//	}
//	remote
//}
//
//type SimpleRemote struct {
//	Remote
//}
//
//func (s *SimpleRemote) MaxVolume() {
//	s.Remote.DeviceInstance.SetVolume(100)
//}
//
//func (s *SimpleRemote) UnVolume() {
//	s.Remote.DeviceInstance.SetVolume(0)
//}
//
//func (s *SimpleRemote) TurnOff() {
//	s.Remote.DeviceInstance.Disable()
//}
//
//func main() {
//	windows := Windows{
//		Devise{
//			DeviceType: windows,
//		},
//	}
//	remote := SimpleRemote{
//		Remote{
//			DeviceInstance: windows,
//		},
//	}
//	remote.MaxVolume()
//	fmt.Println(windows.volume)
//
//}
