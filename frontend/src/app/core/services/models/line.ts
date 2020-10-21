export interface Line {
    id?: number;
    line: number;
    level: string;
    logName: string;
    msg: string;
    timestamp: string | Date;
}
