<?php

namespace App\Resources;

abstract class BaseResource
{
    abstract public function arrayToDto(array $item);
    abstract public function dtoToArray($item): array;
    abstract public function dtoToVo($item);

    public function arrayDtoToVoItens(null|array $itens): array
    {
        if ( ! $itens) {
            return array();
        }
        $itensResourced = array();
        foreach ($itens as $item) {
            $itensResourced[] = $this->dtoToVo($item);
        }
        return $itensResourced;
    }

    public function arrayToDtoItens(null|array $itens): array
    {
        if ( ! $itens) {
            return array();
        }
        $itensResourced = array();
        foreach ($itens as $item) {
            $itensResourced[] = $this->arrayToDto($item);
        }
        return $itensResourced;
    }
}