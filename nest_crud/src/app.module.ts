import { Module } from '@nestjs/common';
import { AppController } from './app.controller';
import { AppService } from './app.service';
import { ConfigModule } from '@nestjs/config';
import { MongooseModule } from '@nestjs/mongoose';
import { ProductModule } from './products/products.module';
@Module({
  imports: [
    ConfigModule.forRoot({ envFilePath: '.env.local', isGlobal: true }),
    MongooseModule.forRoot(process.env.MONGODB_URI),
    ProductModule,
  ],
  controllers: [AppController],
  providers: [AppService],
})
export class AppModule {}
