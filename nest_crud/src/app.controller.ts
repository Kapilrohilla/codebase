import { Controller, Get } from '@nestjs/common';
import { AppService } from './app.service';
@Controller()
export class AppController {
  constructor(private readonly appService: AppService) {}

  @Get()
  getHello(): string {
    // console.log(this.configService.get<string>('MONGODB_URI'));
    return this.appService.getHello();
  }
}
