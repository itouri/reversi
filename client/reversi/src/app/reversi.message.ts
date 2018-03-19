export class ReversiMessage {

  funcName: string;
  body: string;
  systemFlag: boolean;

  constructor(
    funcName: string,
    body: string,
    systemFlag: boolean
  ) {
    this.funcName = funcName;
    this.body = body;
    this.systemFlag = systemFlag;
  }

}
