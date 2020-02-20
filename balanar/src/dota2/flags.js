// https://wiki.teamfortress.com/wiki/WebAPI/GetMatchDetails#Player_Slot
export class PlayerSlot {
  static POS_MASK = 0b111;

  constructor(slot) {
    this.slot = slot || 0;
  }

  get teamNumber() {
    return this.slot >> 7;
  }

  get isRadiant() {
    return this.teamNumber === 0;
  }

  get isDire() {
    return this.teamNumber === 1;
  }

  get position() {
    return this.slot & this.POS_MASK;
  }

  get index() {
    return this.teamNumber * 5 + this.position;
  }
}
