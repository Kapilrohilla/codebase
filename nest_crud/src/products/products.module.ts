import { Module } from '@nestjs/common';
import { ProductController } from './products.controller';
import { ProductService } from './products.service';
import { MongooseModule } from '@nestjs/mongoose';
import { ProductSchema } from '../Schemas/product.schema';
@Module({
  imports: [
    MongooseModule.forFeature([{ name: 'products', schema: ProductSchema }]),
  ],
  controllers: [ProductController],
  providers: [ProductService],
})
export class ProductModule {}
