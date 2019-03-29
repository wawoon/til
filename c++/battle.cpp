#include <iostream>

class Battler {
 public:
  int battle_power;
  Battler(int base_power) { battle_power = base_power; };

  int battle(Battler other) {
    if (battle_power > other.battle_power) {
      return 1;
    } else if (battle_power < other.battle_power) {
      return -1;
    } else {
      return 0;
    };

    int battle_power;
  };
};

int main() {
  auto a = Battler(10);
  auto b = Battler(20);

  std::cout << "Hello, world" << std::endl;
  std::cout << a.battle(b) << std::endl;
  return 0;
}
