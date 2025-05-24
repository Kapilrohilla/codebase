import { Body, Controller, Get, Post } from '@nestjs/common';
import { ReportRequestService } from './report-request.service';
import { CreateReportDto } from './dto';

@Controller('report-request')
export class ReportRequestController {
  constructor(private readonly reportRequestService: ReportRequestService) {}

  @Get('/')
  async getReportRequest() {
    return this.reportRequestService.getReportRequest();
  }

  @Post('/')
  async createReportRequest(@Body() createReportDto: CreateReportDto) {
    try {
      const result =
        await this.reportRequestService.createReportRequest(createReportDto);

      console.log('record created successfully');

      return { success: true, data: result };
    } catch (err: unknown) {
      console.error('Error creating report request:', err);
      return {
        success: false,
        error: err instanceof Error ? err.message : 'Unknown error',
      };
    }
  }
}
