package main

type WeaponStrategy interface {
	UseWeapon()
}

type AK47 struct {
}

func (t *AK47) UseWeapon() {
	println("使用AK47进行战斗")
}

type Knife struct {
}

func (t *Knife) UseWeapon() {
	println("使用匕首进行战斗")
}

type Hero struct {
	weapon WeaponStrategy
}

func (t *Hero) SetWeaponStrategy(strategy WeaponStrategy) *Hero {
	t.weapon = strategy
	return t
}

func (t *Hero) Fight() {
	t.weapon.UseWeapon()
}

func main() {
	ak := new(AK47)
	knife := new(Knife)

	new(Hero).SetWeaponStrategy(ak).Fight()
	new(Hero).SetWeaponStrategy(knife).Fight()
}
