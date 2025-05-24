import { Module } from '@nestjs/common';
import { AppController } from './app.controller';
import { AppService } from './app.service';
import { ConfigModule } from '@nestjs/config';
import { ReportRequestModule } from './report-request/report-request.module';
import { DatabaseModule } from './database/database.module';

@Module({
  imports: [ConfigModule.forRoot(), DatabaseModule, ReportRequestModule],
  controllers: [AppController],
  providers: [AppService],
})
export class AppModule {}
