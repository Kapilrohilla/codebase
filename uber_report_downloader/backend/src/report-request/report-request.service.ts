import { Inject, Injectable } from '@nestjs/common';
import { Model } from 'mongoose';
import { CreateReportDto } from './dto';
import {
  ReportRequest,
  ReportRequestDocument,
} from '../schema/report-request.schema';

@Injectable()
export class ReportRequestService {
  constructor(
    @Inject('REPORT_REQUEST_MODEL')
    private reportRequestModel: Model<ReportRequestDocument>,
  ) {}

  async getReportRequest(): Promise<ReportRequest[]> {
    return this.reportRequestModel.find().exec();
  }

  async createReportRequest(props: CreateReportDto): Promise<ReportRequest> {
    const newReportRequest = new this.reportRequestModel({
      fromDate: props.from,
      toDate: props.to,
      reportType: props.reportType,
      organizations: props.organizations,
      status: 'pending',
    });

    const savedRequest = await newReportRequest.save();

    // Queue the report generation process
    void this.pushReportGenToQueue({
      fromDate: props.from,
      toDate: props.to,
      reportType: props.reportType,
      orgs: props.organizations,
    });

    return savedRequest;
  }

  async pushReportGenToQueue(args: IReportRequest): Promise<boolean> {
    // Here you would implement actual queue logic
    console.log('Queuing report generation job:', args);
    // For example, using a queue service like Bull, RabbitMQ, etc.
    await Promise.resolve(); // Placeholder for async operation
    return true;
  }
}

type IReportRequest = {
  fromDate: Date;
  toDate: Date;
  reportType: string;
  orgs: string[];
};
