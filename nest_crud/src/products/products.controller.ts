import {
  Body,
  Controller,
  Delete,
  Get,
  Param,
  Post,
  Put,
} from '@nestjs/common';
import { ProductService } from './products.service';
import { Product } from './interface/products.interface';
// eslint-disable-next-line @typescript-eslint/no-unused-vars
import { createProductDto } from './dto/createProduct.dto';
// eslint-disable-next-line @typescript-eslint/no-unused-vars
import { updateProductDto } from './dto/updateProduct.dto';
@Controller('products')
export class ProductController {
  constructor(private productService: ProductService) {}
  @Get()
  getProduct(): Promise<Product[]> {
    return this.productService.getProducts();
  }

  @Get(':id')
  getSpecificProduct(@Param('id') id: string): Promise<Product> {
    return this.productService.getSpecificProducts(id);
  }

  @Post()
  createProduct(@Body() createProductDto: createProductDto): Promise<Product> {
    return this.productService.createProducts(createProductDto);
  }

  @Delete(':id')
  deleteProduct(@Param('id') id: string): Promise<Product> {
    return this.productService.deleteProducts(id);
  }

  @Put(':id')
  updateProduct(
    @Param('id') id,
    @Body() updateProductDto: updateProductDto,
  ): Promise<Product> {
    return this.productService.updateProducts(id, updateProductDto);
  }
}
