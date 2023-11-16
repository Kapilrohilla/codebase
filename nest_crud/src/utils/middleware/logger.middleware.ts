import { NestMiddleware, Injectable } from '@nestjs/common';
import { Request, Response, NextFunction } from 'express';
@Injectable()
export class LoggerMiddleware implements NestMiddleware {
  use(req: Request, res: Response, next: NextFunction) {
    res.on('finish', () => {
      const outputString = `${req.method} ${req.url} ${
        res.statusCode
      } ${JSON.stringify(req.body)}`;
      console.log(outputString);
    });
    next();
  }
}
