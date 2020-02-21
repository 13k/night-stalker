// https://wiki.teamfortress.com/wiki/WebAPI/GetMatchDetails#Player_Slot
export class PlayerSlot {
  constructor(slot) {
    this.slot = slot || 0;
    this.team = this.slot >> 7;
    this.position = this.slot & 0b111;
    this.index = this.team * 5 + this.position;
    this.isRadiant = this.team === 0;
    this.isDire = this.team === 1;
  }
}
