import { Module } from '@nestjs/common';
import { ReportRequestController } from './report-request.controller';
import { ReportRequestService } from './report-request.service';
import { DatabaseModule } from '../database/database.module';
import { reportRequestProviders } from './report-request.providers';

@Module({
  imports: [DatabaseModule],
  controllers: [ReportRequestController],
  providers: [...reportRequestProviders, ReportRequestService],
})
export class ReportRequestModule {}
